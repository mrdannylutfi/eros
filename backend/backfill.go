// Psuedocode logic loop
rows, _ := db.Query("SELECT id, img_blob, mime_type FROM images WHERE image_url IS NULL")
for rows.Next() {
    // 1. Read BLOB bytes
    // 2. Upload to S3/Vercel Blob using their SDK
    // 3. UPDATE images SET image_url = $1 WHERE id = $2
}
