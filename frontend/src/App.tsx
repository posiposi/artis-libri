import "./App.css";
import BookTable from "./components/BookTable";
import GetRecommendedBooksButton from "./components/GetRecommendedBooksButton";

function App() {
  return (
    <>
      <h1>Artis Libri</h1>
      <div className="card">
        <GetRecommendedBooksButton />
        <BookTable />
      </div>
    </>
  );
}

export default App;
