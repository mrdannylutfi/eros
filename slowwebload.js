import Image from 'next/image'

export default function Hero() {
  return (
    <Image 
      src="/hero.jpg" 
      alt="Hero Banner" 
      width={1200} 
      height={600} 
      priority={true} // High fetchpriority; bypasses default lazy-loading
      sizes="(max-width: 768px) 100vw, 1200px"
    />
  )
}

/*

 2. Fix Slow Visual Rendering (LCP)

LCP delays occur because the browser takes too long to discover, download, or paint the largest visible block on the screen (usually a hero image or main title text).
• Apply Priority to Hero Images: If your LCP element is an image, the fastest fix is marking it with priority={true} within the Next.js Image component. This instructs Next.js to append a high-priority preload link tag to your document header, ensuring the browser fetches it immediately rather than waiting to discover it down the DOM tree.

*/
