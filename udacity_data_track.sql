Select id
From orders
Where ( (poster_qty > '4000' OR gloss_qty > '4000') );

Select *
From orders
where standard_qty = '0'
AND (poster_qty > '1000' OR gloss_qty > '1000');

Select *
FROM accounts
WHERE name LIKE 'C%'
OR name LIKE 'W%'
AND( primary_poc LIKE ('%ana%') OR primary_poc LIKE ('%Ana%'));
AND primary_poc NOT LIKE ('%eana%');
