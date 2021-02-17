import React from "react";
import { ReactComponent as CartIcon } from "../../assets/shopping-cart.svg";
import { useSelector } from "react-redux";
import styles from "./index.module.css";
import { useHistory } from 'react-router-dom'

export default function Cart(props) {
  const items = useSelector((state) => state.user.cart);
  const count = (items || []).length;
  const history = useHistory()

  const goToCart = () => history.push('/cart')

  return (
    <div className={styles.container} onClick={goToCart}>
      {count ? <div className={styles.count}>{count}</div> : <div />}
      <CartIcon className={styles.icon} />
    </div>
  );
}
