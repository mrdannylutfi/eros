ALTER TABLE images DROP COLUMN img_blob;
VACUUM FULL images;
