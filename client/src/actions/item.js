export const ADD_TO_CART = "ADD_TO_CART"
export const addItemToCart = (item) => {
  return {
    type: ADD_TO_CART,
    payload: item
  }
}

export const ADD_TO_WISHLIST = "ADD_TO_WISHLIST"
export const wishList = (item) => {
  return {
    type: ADD_TO_WISHLIST,
    payload: item
  }
}

export const REMOVE_FROM_CART = "REMOVE_FROM_CART"
export const removeFromCart = (itemID) => {
  return {
    type: REMOVE_FROM_CART,
    payload: itemID
  }
}

export const REMOVE_FROM_WISHLIST = "REMOVE_FROM_WISHLIST"
export const removeFromWishList = (itemID) => {
  return {
    type: REMOVE_FROM_WISHLIST,
    payload: itemID
  }
}