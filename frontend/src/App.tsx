import "./App.css";
import BookRegisterButton from "./components/BookRegisterButton";
import BookTable from "./components/BookTable";

function App() {
  return (
    <>
      <h1>Artis Libri</h1>
      <div className="book_register_btn">
        <BookRegisterButton />
      </div>
      <div className="card">
        <BookTable />
      </div>
    </>
  );
}

export default App;
