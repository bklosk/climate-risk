package main

const Query = `WITH boundingbox AS(
	SELECT
	  ST_MakeEnvelope(
		%(xmin)s,
		%(ymin)s,
		%(xmax)s,
		%(ymax)s,
		3857
	  ) AS geometry
  ),
  mvtgeom AS (
	SELECT
	  NAME,
	  ST_AsMVTGeom(
		ST_Transform(sourceTable.geom, 3857),
		boundingbox.geometry
	  )
	FROM
	  "netcdf:test_data" sourceTable,
	  boundingbox
	WHERE
	  ST_Intersects(
		ST_Transform(boundingbox.geometry, 4326),
		sourceTable.geom
	  )
  )
  SELECT
	ST_AsMVT(mvtgeom.*)
  FROM
	mvtgeom;`