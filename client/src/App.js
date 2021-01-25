import logo from "./logo.svg";
import "./App.css";
import Button from "./components/button";
import Input from "./components/input";
import Card from "./components/card";
import Dropdown from "./components/dropdown";
import AppBar from "./components/appbar";

function App() {
  return (
    <>
      <AppBar>Yo</AppBar>
      <Card>
        <Button>test</Button>
        <Input />
        <Dropdown />
      </Card>
    </>
  );
}

export default App;
