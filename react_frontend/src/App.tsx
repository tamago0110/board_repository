import React from "react";
import styles from "./App.module.css";
import Navbar from "./features/Navbar";
import Main from "./features/Main";

const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <div className={styles.container}>
        <Main />
      </div>
    </>
  );
};

export default App;
