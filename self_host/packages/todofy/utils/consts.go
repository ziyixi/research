package utils

const (
	KeyGRPCClients = "grpcClients"

	DefaultpromptToSummaryEmail string = `Could you please provide a concise and comprehensive summary of the given email? The summary should capture the main points and key details of the text while conveying the author's intended meaning accurately. Please ensure that the summary is well-organized and easy to read, with clear headings and subheadings to guide the reader through each section. The length of the summary should be appropriate to capture the main points and key details of the text, without including unnecessary information or becoming overly long. 
	
	IMPORTANT: Please do not write something like "OK, this is my summary". Just start with the summary.
	IMPORTANT: Try to follow markdown formatting as much as possible.

	The email content you are to summarize is as follows:`
)
