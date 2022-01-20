package gqltest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/machinebox/graphql"
)

type Recorder struct {
	Addr  string
	Token string

	client *graphql.Client
	out    []byte
}

func (r *Recorder) Query(title, qry string, args ...interface{}) []byte {
	if r.client == nil {
		r.client = graphql.NewClient(r.Addr)
	}

	qry = formatPayload(qry, args...)

	res := r.execute(title, qry, true)

	return res
}

func (r *Recorder) QuerySilent(title, qry string, args ...interface{}) []byte {
	if r.client == nil {
		r.client = graphql.NewClient(r.Addr)
	}

	qry = formatPayload(qry, args...)

	res := r.execute(title, qry, false)

	return res
}

func (r *Recorder) Mutate(title, qry string, args ...interface{}) []byte {
	if r.client == nil {
		r.client = graphql.NewClient(r.Addr)
	}

	qry = formatPayload(qry, args...)

	res := r.execute(title, qry, true)

	return res
}

func (r *Recorder) MutateSilent(title, qry string, args ...interface{}) []byte {
	if r.client == nil {
		r.client = graphql.NewClient(r.Addr)
	}

	qry = formatPayload(qry, args...)

	res := r.execute(title, qry, false)

	return res
}

func (r *Recorder) Comment(cmt string, args ...interface{}) {
	cmtB := []byte(fmt.Sprintf(cmt, args...))
	cmtB = bytes.ReplaceAll(cmtB, []byte("\n"), []byte("\n# "))

	r.out = append(r.out, '#', ' ')
	r.out = append(r.out, cmtB...)
	r.out = append(r.out, '\n')
	r.out = append(r.out, '\n')
}

func (r *Recorder) execute(title, qry string, output bool) []byte {
	req := graphql.NewRequest(qry)

	if r.Token != "" {
		req.Header.Add("Cookie", "session="+r.Token)
	}

	res := json.RawMessage{}
	resErr := r.client.Run(context.Background(), req, &res)

	var (
		enc []byte
		err error
	)

	if resErr == nil {
		enc, err = json.MarshalIndent(res, "", "\t")
		if err != nil {
			panic(err)
		}
	}

	if output {
		r.out = append(r.out, buildStepOutput(title, qry, enc, resErr)...)
	}

	return res
}

func (r *Recorder) Summary() []byte {
	return r.out
}

func buildStepOutput(title, req string, res []byte, resErr error) []byte {
	out := []byte{}
	line := strings.Repeat("─", len(title))

	out = append(out, "┌─"+line+"─┐\n"...)
	out = append(out, "│ "+title+" │\n"...)
	out = append(out, "└─"+line+"─┘"...)
	out = append(out, '\n')
	out = append(out, '\n')
	out = append(out, req...)
	out = append(out, '\n')
	out = append(out, '\n')

	if resErr != nil {
		out = append(out, fmt.Sprintf("Failure: %s", resErr)...)
	} else {
		out = append(out, res...)
	}

	out = append(out, '\n')
	out = append(out, '\n')

	return out
}

func formatPayload(orig string, args ...interface{}) string {
	// Remove the extra tabs at the
	// beginning of each line.
	start := "\n"
	fit := true

	for fit {
		start += "\t"
		fit = strings.HasPrefix(orig, start+"\t")
	}

	orig = strings.ReplaceAll(orig, start, "\n")
	orig = strings.TrimSpace(orig)
	orig = fmt.Sprintf(orig, args...)

	return orig
}
