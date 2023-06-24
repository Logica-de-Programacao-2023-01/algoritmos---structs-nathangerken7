package main

type Musica struct {
	titulo  string
	artista string
	duracao int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func encontrarMusica(playlists []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado, playlist)
			}
		}
	}
	return resultado
}
