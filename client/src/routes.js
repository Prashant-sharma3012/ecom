import React from "react";
import { Switch, Route } from "react-router-dom";
import Home from "./pages/home";
import Cart from "./pages/cart";
import ItemDetail from "./pages/itemDetail";


export default function Routes() {
  return (
    <>
      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
        <Route path="/cart">
          <Cart />
        </Route>
        <Route path="/item/:itemId">
          <ItemDetail />
        </Route>
      </Switch>
    </>
  );
}
