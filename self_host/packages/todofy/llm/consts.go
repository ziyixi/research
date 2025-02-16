package main

import pb "github.com/ziyixi/monorepo/self_host/packages/todofy/proto"

var (
	llmModelNames = map[pb.Model]string{
		pb.Model_MODEL_GEMINI_2_0_PRO_EXP_02_05: "gemini-2.0-pro-exp-02-05",
		pb.Model_MODEL_GEMINI_1_5_PRO:           "gemini-1.5-pro",
		pb.Model_MODEL_GEMINI_2_0_FLASH:         "gemini-2.0-flash",
		pb.Model_MODEL_GEMINI_1_5_FLASH:         "gemini-1.5-flash",
	}
	llmModelPriority = []pb.Model{
		pb.Model_MODEL_GEMINI_2_0_PRO_EXP_02_05,
		pb.Model_MODEL_GEMINI_1_5_PRO,
		pb.Model_MODEL_GEMINI_2_0_FLASH,
		pb.Model_MODEL_GEMINI_1_5_FLASH,
	}
	supportedModelFamily = []pb.ModelFamily{
		pb.ModelFamily_MODEL_FAMILY_GEMINI,
	}
)

const (
	DefaultpromptToSummaryEmail string = `Could you please provide a concise and comprehensive summary of the given email? The summary should capture the main points and key details of the text while conveying the author's intended meaning accurately. Please ensure that the summary is well-organized and easy to read, with clear headings and subheadings to guide the reader through each section. The length of the summary should be appropriate to capture the main points and key details of the text, without including unnecessary information or becoming overly long. 
	
	IMPORTANT: Please do not write something like "OK, this is my summary". Just start with the summary.
	IMPORTANT: Try to follow markdown formatting as much as possible.

	The email content you are to summarize is as follows:`

	tokenLimit int32 = 1048576 // 10k tokens
)
