import React from "react";
import styles from "./index.module.css";
import { ReactComponent as Night } from "../../assets/night-mode.svg";
import { ReactComponent as Day } from "../../assets/sunny-day.svg";
import Cart from '../cart'
import { useTheme } from "../../hooks/useTheme";

export default function AppBar(props) {
  const [theme, toggleTheme] = useTheme();

  return (
    <div className={styles.container}>
      <div className={styles.brand}>Khareedari</div>
      <div className={styles.rightSide}>
        <div className={styles.iconContainer} onClick={toggleTheme}>
          {theme === "light" ? <Night className={styles.icon} /> : <Day  className={styles.icon} />}
        </div>
        <div className={styles.user}>Luffy</div>
        <div className={styles.iconContainer}>
          <Cart />
        </div>
      </div>
    </div>
  );
}
