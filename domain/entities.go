package domain

type Anime struct {
	ID          int     `json:"id" ksql:"id"`
	Name        *string `json:"name" ksql:"name"`
	Description *string `json:"description" ksql:"description"`
	NumEpisodes *int    `json:"num_episodes" ksql:"num_episodes"`
	Img         *string `json:"img" ksql:"img"`
}
