import "./App.css";
import BookRegisterDialog from "./components/BookRegisterDialog";
import BookTable from "./components/BookTable";

function App() {
  return (
    <>
      <h1>Artis Libri</h1>
      <div className="book_register_btn">
        <BookRegisterDialog />
      </div>
      <div className="card">
        <BookTable />
      </div>
    </>
  );
}

export default App;
