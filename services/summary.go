package services

import (
	"context"
	"fmt"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/vertexai/genai"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

func GetSummaries(uid string) (structs.GetSummariesRes, error) {
    ctx := context.Background()
    client, err := utils.NewFirestoreClient()
    if err != nil {
        return structs.GetSummariesRes{}, err
    }
    defer client.Close()

    iter := client.Collection("summaries").Where("uid", "==", uid).OrderBy("timestamp", firestore.Desc).Documents(ctx)
    var summaries []structs.Summary

    // イテレーターを使用してドキュメントを取得
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return structs.GetSummariesRes{}, err
        }

        var summary structs.Summary
        if err := doc.DataTo(&summary); err != nil {
            return structs.GetSummariesRes{}, err
        }
        summary.SummaryId = doc.Ref.ID
        summaries = append(summaries, summary)
    }

	// サマリーが見つからない場合は空の配列を返す
	if len(summaries) == 0 {
		return structs.GetSummariesRes{
			Summaries: []structs.Summary{},
		}, nil
	}

    return structs.GetSummariesRes{
        Summaries: summaries,
    }, nil
}

func GetSelectedSummary(summaryId string) (structs.Summary, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.Summary{}, err
	}
	defer client.Close()

	docRef := client.Collection("summaries").Doc(summaryId)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return structs.Summary{}, err
	}

	var summary structs.Summary
	if err := doc.DataTo(&summary); err != nil {
		return structs.Summary{}, err
	}
	
	summary.SummaryId = doc.Ref.ID

	return summary, nil
}

func RequestSummary(uid, content string) (structs.SummaryRequestRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.SummaryRequestRes{}, err
	}
	defer client.Close()

	summaryId := uuid.New().String()

	// AI API へのリクエストと要約結果を取得する処理（ここにAI呼び出しのコードを追加）
	summaryContent, err := generateSummaryAI(content) // generateSummaryAI関数はAIで要約を作成する関数とします
	if err != nil {
		return structs.SummaryRequestRes{}, err
	}

	// Firestore に保存するデータ構造を作成
	docRef := client.Collection("summaries").Doc(summaryId)
	_, err = docRef.Set(ctx, map[string]interface{}{
		"uid":     uid,
		"content": content,
		"summary": summaryContent,
		"timestamp": time.Now(),
	})

	if err != nil {
		return structs.SummaryRequestRes{}, err
	}

	// 保存が成功したら、summaryId をレスポンスとして返す
	response := structs.SummaryRequestRes{
		SummaryId: summaryId,
	}
	return response, nil
}

func generateSummaryAI(content string) (string, error) {
    ctx := context.Background()
    client, err := utils.CreateVertexAIClient()
    if err != nil {
        return "", fmt.Errorf("failed to create Vertex AI client: %v", err)
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-pro")
    
    prompt := `以下の内容を要約し、JSON形式で返してください:
    {
        "title": "内容を端的に表すタイトル（30文字以内）",
        "summary": "内容の要約（400文字以内）"
    }

    内容:
    ` + content

    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", fmt.Errorf("generate content error: %v", err)
    }

    return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}