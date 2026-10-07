
# Vercel Platform Fixes (App & Pages):

## Lean Out middleware.ts: 
Vercel runs middleware globally at the Edge. 
If your middleware connects to an external database, 
executes complex crypto functions, or matches all static paths, 
it delays the response. Add a strict config.matcher array to exclude static files, 
images, and fonts from running through your middleware.

