import styles from "@/styles/Landing.module.css";
import Link from "next/link";

export default function ChamaCard(props: any): JSX.Element {
  return (
    <Link href={`/chama/${props.id}`}>
      <div className={styles.chama_card}>
        <div className={styles.chama_image}></div>
        <div className={styles.left}>
          <div className={styles.chama_name}><b>{props.name}</b></div>
          <div className={styles.chama_desc}>{props.desc}</div>
          <div className={styles.chama_creator}>@{props.creator}</div>
        </div>
      </div>
    </Link>
  );
}
