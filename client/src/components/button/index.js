import React from "react";
import styles from "./index.module.css";

export default function Button(props) {
  return (
    <button className={styles.btn} onClick={props.onClick}>
      {props.children}
    </button>
  );
}
