package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const wikimediaSearch = "https://en.wikipedia.org/w/rest.php/v1/search/page?limit=1&q="
const wikisummary = "https://en.wikipedia.org/api/rest_v1/page/summary/"

type WikiSumJSON struct {
	Title     string
	Extract   string
	Thumbnail struct {
		Source string
	}
}

type WikiSearchJSON struct {
	Key string
}

type WikiSearchResult struct {
	Pages []*WikiSearchJSON
}

// Cada função procura fará 2 requests caso o primeiro seja bem sucedido
func (p *Pessoa) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(p.nome)
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", p.nome, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (p *Pessoa) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", p.nome, err)
	}

	p.URLimage = summary.Thumbnail.Source
	p.descricao = summary.Extract
	p.titulo = summary.Title
	return nil
}

func (loc *Local) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(loc.nome)
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", loc.nome, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (loc *Local) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", loc.nome, err)
	}

	loc.titulo = summary.Title
	loc.URLimage = summary.Thumbnail.Source
	loc.descricao = summary.Extract
	return nil
}

func (aero *Aeroporto) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(aero.codigo + "_airport")
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", aero.codigo, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (aero *Aeroporto) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", aero.codigo, err)
	}

	aero.URLimage = summary.Thumbnail.Source
	aero.descricao = summary.Extract
	aero.titulo = summary.Title
	return nil
}
