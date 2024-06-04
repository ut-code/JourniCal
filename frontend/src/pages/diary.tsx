import { useState, useEffect } from "react";
import DiaryEntry from "../components/DiaryEntry";

type Entry = {
  date: Date;
  title: string;
  content: string;
};

function Diary() {
  const [showPopup, setShowPopup] = useState(false);
  const [editingIndex, setEditingIndex] = useState<number | null>(null);
  const [newEntry, setNewEntry] = useState({ date: "", title: "", content: "" });
  const [diaries, setDiaries] = useState<Entry[] | null>(null);

  // 日記を取得してdiariesにセットする
  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:3000/api/diaries");
      const data = await response.json();

      // date を Date 型に変換
      const entries: Entry[] = data.map((entry: Omit<Entry, "date"> & { date: string }) => ({
        date: new Date(entry.date),
        title: entry.title,
        content: entry.content,
      }));

      // 日付順にソートする
      const sortedEntries = entries.sort(
        (a, b) => new Date(a.date).getTime() - new Date(b.date).getTime()
      );

      setDiaries(sortedEntries);
    })();
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    if (name === 'date') {
      // 日付が入力されていない場合、今日の日付を設定する
      const currentDate = value ? value : new Date().toISOString().slice(0, 10);
      setNewEntry(prevEntry => ({
        ...prevEntry,
        date: currentDate
      }));
    } else {
      setNewEntry(prevEntry => ({
        ...prevEntry,
        [name]: value
      }));
    }
  };

  const handleAddEntry = () => {
    if (!diaries) return;
    const dateObject = new Date(newEntry.date);
    
    if (editingIndex !== null) {
      const updatedDiaries = [...diaries];
      updatedDiaries[editingIndex] = { date: dateObject, title: newEntry.title, content: newEntry.content };
      setDiaries(updatedDiaries);
      setEditingIndex(null);
    } else {
      setDiaries([...diaries, { date: dateObject, title: newEntry.title, content: newEntry.content }]);
    }
    setShowPopup(false);
    setNewEntry({ date: "", title: "", content: "" });
  };

  const handleEditEntry = (index: number) => {
    if (!diaries) return;
    const entryToEdit = diaries[index];
    // 日付を文字列に変換
    const dateString = entryToEdit.date.toISOString().substring(0, 10);
    setNewEntry({
      date: dateString,
      title: entryToEdit.title,
      content: entryToEdit.content
    });
    setEditingIndex(index);
    setShowPopup(true);
  };
  

  const handleCancelEdit = () => {
    setEditingIndex(null);
    setShowPopup(false);
    setNewEntry({ date: "", title: "", content: "" });
  };

  const handleSaveEntry = () => {
    handleAddEntry();
  };

  return (
    <div className="diary-app">
      <div className="diary-entries">
        {diaries && diaries.map((diary, index) => (
          <div key={index}>
            <DiaryEntry {...diary} onEdit={() => handleEditEntry(index)} />
          </div>
        ))}
      </div>
      {showPopup && (
        <div style={{
          position: 'fixed',
          top: 0,
          left: 0,
          width: '100%',
          height: '100%',
          backgroundColor: 'rgba(0, 0, 0, 0.5)',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}>
          <div style={{
            backgroundColor: 'white',
            padding: '20px',
            borderRadius: '5px',
            display: 'flex',
            flexDirection: 'column',
          }}>
            <input type="date" name="date" value={newEntry.date} onChange={handleInputChange} />
            <input type="text" name="title" placeholder="Title" value={newEntry.title} onChange={handleInputChange} />
            <textarea name="content" placeholder="Content" value={newEntry.content} onChange={handleInputChange} />
            <button onClick={handleCancelEdit}>Cancel</button>
            <button onClick={handleSaveEntry}>{editingIndex !== null ? "Update Entry" : "Add Entry"}</button>
          </div>
        </div>
      )}
      <button onClick={() => setShowPopup(true)}>+</button>
    </div>
  );
}

export default Diary;
