import Image from 'next/image'
import Link from 'next/link'
import styles from '@/styles/Home.module.css'


function Logo() : JSX.Element {
  return (
    <Link href="/">
        <div className={styles.logo}>
          <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
        </div>
        </Link>
  )
}

export default Logo
