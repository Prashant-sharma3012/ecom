import React from "react";
import Rating from "../rating"
import styles from "./index.module.css"
import Button from "../button"
import { useDispatch } from 'react-redux'
import { addItemToCart, wishList } from '../../actions/item'

const defaultItem = {
  image: "//i.dell.com/is/image/DellContent//content/dam/global-asset-library/Products/Notebooks/Inspiron/15_5508_non-touch/in5508nt_cnb_00055lf110_gr.psd?fmt=pjpg&pscan=auto&scl=1&hei=402&wid=668&qlt=95,0&resMode=sharp2&op_usm=1.75,0.3,2,0&size=668,402",
  intro: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
  seller: "Acer India",
  price: "$5000"
}

export default function Item({item = defaultItem, ...props}) {
  const dispatch = useDispatch();
  const addToCart = () => dispatch(addItemToCart(defaultItem))
  const addToWishList = () => dispatch(wishList(defaultItem))
  
  return (
    <div className={styles.container}>
      <div className={styles.banner}>
        <img className={styles.itemImage} src={item.image} />
      </div>
      <div className={styles.intro}>
        {item.intro}
      </div>
      <div className={styles.seller}>
        {item.seller}
      </div>
      <div className={styles.footer}>
        <div className={styles.rating}>
          <Rating />
        </div>
        <div className={styles.price}>
          {item.price}
        </div>
      </div>
      <div className={styles.actionBtn}>
        <Button onClick={addToWishList}>Wishlist</Button>
        <Button onClick={addToCart}>Add to cart</Button>
      </div>
    </div>
  );
}
