export async function getStaticProps() {
  return {
    props: { data },
    revalidate: 60, // Regenerates in the background at most once every 60 seconds
  }
}
<div>
  {/* This is a comment inside JSX :

## Use Incremental Static Regeneration (ISR): 
If you need to update data periodically without a full `rebuild`, add the revalidate `prop` to your static props:
*/
</div>
