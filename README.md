# climate-risk
Webapp that uses climate models to predict how a user's microclimate will change

/src holds a React / Tailwind app using D3 to visualize user outcomes\
/web-service holds a Go (Gin) RESTful API: street address -> climate time-series\
/web-service/utilities holds data modification utilities\
/research holds climate models, sources, and thoughts

React app (digitalocean) <-> Gin API (digitalocean) <-> Postgres(postGIS) (AWS)