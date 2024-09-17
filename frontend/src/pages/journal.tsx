import JournalEntry from "../components/JournalEntry";
import useJournal from "../hooks/useJournal";

function Journal() {
  const { journals, isLoading, error } = useJournal();
  return (
    <div className="journal-app">
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error: {error.message}</div>
      ) : (
        <div className="journal-entries">
          {journals?.map((journal, index) => (
            <JournalEntry key={index} journal={journal} />
          ))}
        </div>
      )}
    </div>
  );
}

export default Journal;
