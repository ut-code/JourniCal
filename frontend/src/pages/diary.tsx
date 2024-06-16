import DiaryEntry from "../components/DiaryEntry";

interface Entry {
  date: string;
  title: string;
  content: string;
}

async function fetchEntries(): Promise<Entry[]> {
  const response = await fetch("http://localhost:3000/api/diaries");
  const data = (await response.json()) as Entry[];
  console.log(data);
  return data;
}

// const diaries = await fetchEntries();
const diaries: Entry[] = [];

function Diary() {
  return (
    <div className="diary-app">
      <div className="diary-entries">
        {diaries.map((diary, index) => (
          <DiaryEntry key={index} {...diary} />
        ))}
      </div>
    </div>
  );
}

export default Diary;
