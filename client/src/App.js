import styles from "./App.module.css";
import Routes from "./routes"
import AppBar from './components/appbar'

function App() {
  
  return (
    <div className={styles.main}>
      <div className={styles.appbar}>
        <AppBar />
      </div>
      <div className={styles.content}>
        <Routes />
      </div>
    </div>
  );
}

export default App;
