import Link from "next/link";
import styles from "@/styles/Landing.module.css";

export default function ProposalCard(props: any): JSX.Element {
  return (
    <Link href={`/proposal/${props.id}`}>
      <div className={styles.chama_card}>
        <div className={styles.chama_image}></div>
        <div className={styles.left}>
          <div className={styles.chama_name}>{props.name}</div>
          <div className={styles.chama_desc}>{props.desc}</div>
          <div className={styles.chama_creator}>@{props.creator}</div>
        </div>
      </div>
    </Link>
  );
}
