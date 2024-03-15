	result := db.Debug().Table(`"6PMS_1RUN".run40prohd as run40`).
		Select(`run40."opcode",run40."dealN",run40."N",run40."currencytype",run40."sumtotal",run40."tax",run40."total",run40."suppalias",run40."suppID",run40."exchangerate",run40."driID",sdm60."suppNM",sdm50."peopleNM" as "driNM", run41."openCount"`).
		Joins(`LEFT JOIN "0ADM_1SDM".sdm50staff as sdm50 on sdm50."staffID" = run40."driID"`).
		Joins(`LEFT JOIN "0ADM_1SDM".sdm60supp as sdm60 on sdm60."suppID" = run40."suppID"`).
		Joins(`LEFT JOIN LATERAL (
			SELECT count("opUUID") as "openCount" FROM "6PMS_1RUN".run41protb as run41
			WHERE run41."opUUID" = run40."opUUID" AND run41."lifeF" != 'T'
		) as run41 on true`).
		Where(`run40."lifeF" in ('2','7','T')`).
		Order(`run40."suppID" DESC`).
		Find(&res)