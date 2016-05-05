CREATE MATERIALIZED VIEW data.latency_summary
AS WITH summ as (SELECT sitePK, typePK, time, mean, fifty, ninety
	FROM
	(SELECT sitePK, typePK, time, mean, fifty, ninety, rank()
		OVER ( PARTITION BY sitePK, typePK ORDER BY time DESC) FROM data.latency) s
	WHERE rank = 1)
	SELECT sitePK, typePK, siteID, typeID, geom, time, mean, fifty, ninety,
  CASE WHEN latency_threshold.lower is NULL THEN 0 ELSE latency_threshold.lower END AS "lower",
	CASE WHEN latency_threshold.upper is NULL THEN 0 ELSE latency_threshold.upper END AS "upper"
  from summ
	JOIN data.site USING (sitePK)
	LEFT OUTER JOIN data.latency_threshold USING (sitePK, typePK)
	JOIN data.type using (typePK);

-- UNIQUE index is needed for refresh CONCURRENTLY
CREATE UNIQUE INDEX on data.latency_summary (sitePK, typePk, time);

GRANT SELECT ON data.latency_summary TO mtr_r;

ALTER MATERIALIZED VIEW data.latency_summary OWNER TO mtr_w;