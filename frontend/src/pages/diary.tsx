import DiaryEntry from "../components/DiaryEntry";
import useDiary from "../hooks/useDiary";

function Diary() {
  const { diaries, isLoading, error } = useDiary();
  return (
    <div className="diary-app">
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error: {error.message}</div>
      ) : (
        <div className="diary-entries">
          {diaries?.map((diary, index) => (
            <DiaryEntry key={index} {...diary} />
          ))}
        </div>
      )}
    </div>
  );
}

export default Diary;
