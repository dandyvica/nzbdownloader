// list of RFC3977 response codes to commands
package main

// list of codes here: https://datatracker.ietf.org/doc/html/rfc3977#appendix-C
type ResponseCodeMeaning struct {
	GeneratedBy string // list of commands generating the response code
	Meaning     string // response code description
}

// all response codes
var responseCodes = map[string]ResponseCodeMeaning{
	"100":   ResponseCodeMeaning{GeneratedBy: "HELP", Meaning: "help text follows."},
	"101":   ResponseCodeMeaning{GeneratedBy: "CAPABILITIES", Meaning: "capabilities list follows."},
	"111":   ResponseCodeMeaning{GeneratedBy: "DATE", Meaning: "server date and time."},
	"200":   ResponseCodeMeaning{GeneratedBy: "initial connection, MODE READER", Meaning: "service available, posting allowed."},
	"201":   ResponseCodeMeaning{GeneratedBy: "initial connection, MODE READER", Meaning: "service available, posting prohibited."},
	"205":   ResponseCodeMeaning{GeneratedBy: "QUIT", Meaning: "connection closing (the server immediately closes the connection)."},
	"211":   ResponseCodeMeaning{GeneratedBy: "GROUP", Meaning: "group selected."},
	"211-1": ResponseCodeMeaning{GeneratedBy: "LISTGROUP", Meaning: "article numbers follow."},
	"215":   ResponseCodeMeaning{GeneratedBy: "LIST", Meaning: "information follows."},
	"220":   ResponseCodeMeaning{GeneratedBy: "ARTICLE", Meaning: "article follows."},
	"221":   ResponseCodeMeaning{GeneratedBy: "HEAD", Meaning: "article headers follow."},
	"222":   ResponseCodeMeaning{GeneratedBy: "BODY", Meaning: "article body follows."},
	"223":   ResponseCodeMeaning{GeneratedBy: "LAST, NEXT, STAT", Meaning: "article exists and selected."},
	"224":   ResponseCodeMeaning{GeneratedBy: "OVER", Meaning: "overview information follows."},
	"225":   ResponseCodeMeaning{GeneratedBy: "HDR", Meaning: "headers follow."},
	"230":   ResponseCodeMeaning{GeneratedBy: "NEWNEWS", Meaning: "list of new articles follows."},
	"231":   ResponseCodeMeaning{GeneratedBy: "NEWGROUPS", Meaning: "list of new newsgroups follows."},
	"235":   ResponseCodeMeaning{GeneratedBy: "IHAVE (second stage)", Meaning: "article transferred OK."},
	"240":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "article received OK."},
	"335":   ResponseCodeMeaning{GeneratedBy: "IHAVE (first stage)", Meaning: "send article to be transferred."},
	"340":   ResponseCodeMeaning{GeneratedBy: "POST (first stage)", Meaning: "send article to be posted."},
	"400":   ResponseCodeMeaning{GeneratedBy: "POST (first stage)", Meaning: "service not available or no longer available (the server immediately closes the connection)."},
	"401":   ResponseCodeMeaning{GeneratedBy: "POST (first stage)", Meaning: "the server is in the wrong mode; the indicated capability should be used to change the mode."},
	"403":   ResponseCodeMeaning{GeneratedBy: "POST (first stage)", Meaning: "internal fault or problem preventing action being taken."},
	"411":   ResponseCodeMeaning{GeneratedBy: "GROUP, LISTGROUP", Meaning: "no such newsgroup."},
	"412":   ResponseCodeMeaning{GeneratedBy: "ARTICLE, BODY, GROUP, HDR, HEAD, LAST, LISTGROUP,", Meaning: "no newsgroup selected."},
	"420":   ResponseCodeMeaning{GeneratedBy: "ARTICLE, BODY, HDR, HEAD, LAST, NEXT, OVER, STAT", Meaning: "current article number is invalid."},
	"421":   ResponseCodeMeaning{GeneratedBy: "NEXT", Meaning: "no next article in this group."},
	"422":   ResponseCodeMeaning{GeneratedBy: "LAST", Meaning: "no previous article in this group."},
	"423":   ResponseCodeMeaning{GeneratedBy: "ARTICLE, BODY, HDR, HEAD, OVER, STAT", Meaning: "no article with that number or in that range."},
	"430":   ResponseCodeMeaning{GeneratedBy: "ARTICLE, BODY, HDR, HEAD, OVER, STAT", Meaning: "no article with that message-id."},
	"435":   ResponseCodeMeaning{GeneratedBy: "IHAVE (first stage)", Meaning: "article not wanted."},
	"436":   ResponseCodeMeaning{GeneratedBy: "IHAVE (either stage)", Meaning: "transfer not possible (first stage) or failed (second stage); try again later."},
	"437":   ResponseCodeMeaning{GeneratedBy: "IHAVE (second stage)", Meaning: "transfer rejected; do not retry."},
	"440":   ResponseCodeMeaning{GeneratedBy: "POST (first stage)", Meaning: "posting not permitted."},
	"441":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "posting failed."},
	"480":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "command unavailable until the client has authenticated itself."},
	"483":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "command unavailable until suitable privacy has been"},
	"500":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "unknown command."},
	"501":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "syntax error in command."},
	"503":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "feature not supported."},
	"504":   ResponseCodeMeaning{GeneratedBy: "POST (second stage)", Meaning: "error in base64-encoding [RFC4648] of an argument."},
}
