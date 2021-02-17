import {
  ADD_TO_CART,
  ADD_TO_WISHLIST,
  REMOVE_FROM_CART,
  REMOVE_FROM_WISHLIST,
} from "../actions/item";
import produce from "immer"

const initialState = {
  cart: [],
  wishlist: [],
};

const ItemReducer = produce((draft , action) => {
  switch (action.type) {
    case ADD_TO_CART:
      draft.cart.push(action.payload)
      break;
    case ADD_TO_WISHLIST:
      draft.wishlist.push(action.payload);
      break;
    case REMOVE_FROM_CART:
      draft.cart = draft.cart.filter(item => item.id !== action.payload)
      break;
    case REMOVE_FROM_WISHLIST:
      draft.wishlist = draft.wishlist.filter(item => item.id !== action.payload)
      break;
  }
}, initialState)

export default ItemReducer;
