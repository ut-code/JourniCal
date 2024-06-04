import { useState, useEffect } from "react";
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

const fetchedDiaries = await fetchEntries();

function Diary() {
  const [showPopup, setShowPopup] = useState(false);
  const [editingIndex, setEditingIndex] = useState(null);
  const [newEntry, setNewEntry] = useState({ date: "", title: "", content: "" });
  const [diaries, setDiaries] = useState([
    //例
    {
      date: new Date("2024-03-30"),
      title: "First Entry",
      content: "あいうえお",
    },
  ]);

  //日付順にソートする
  useEffect(() => {
    const sortedDiaries = [...diaries].sort((a, b) => a.date.getTime() - b.date.getTime());
    setDiaries(sortedDiaries);
  }, [diaries]);

  const handleInputChange = (event:any) => {
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

  const handleEditEntry = (index:any) => {
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
        {diaries.map((diary, index) => (
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
