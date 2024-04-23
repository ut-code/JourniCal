import  { useState, useEffect } from "react";
import DiaryEntry from "../components/DiaryEntry";

function DiaryApp() {
  const [showPopup, setShowPopup] = useState(false);
  const [editingIndex, setEditingIndex] = useState(null);
  const [newEntry, setNewEntry] = useState({ date: Date(), title: "", content: "" });
  const [diaries, setDiaries] = useState([
    {
      date: new Date("2024-03-30"),
      title: "First Entry",
      content: "あいうえお",
    },
    // Existing diary entries
  ]);

  // 日記エントリーが変更された時に再ソートする
  useEffect(() => {
    const sortedDiaries = [...diaries].sort((a, b) => a.date.getTime() - b.date.getTime());
    setDiaries(sortedDiaries);
  }, [diaries]);

  const handleInputChange = (event:any) => {
    const { name, value } = event.target;
    setNewEntry(prevEntry => ({
      ...prevEntry,
      [name]: value
    }));
  };

  const handleAddEntry = () => {
    const dateObject = new Date(newEntry.date);
    if (editingIndex !== null) {
      // 編集モードの場合は更新
      const updatedDiaries = [...diaries];
      updatedDiaries[editingIndex] = { date: dateObject, title: newEntry.title, content: newEntry.content };
      setDiaries(updatedDiaries);
      setEditingIndex(null);
    } else {
      // 新規追加
      setDiaries([...diaries, { date: dateObject, title: newEntry.title, content: newEntry.content }]);
    }
    setShowPopup(false);
    setNewEntry({ date: "", title: "", content: "" });
  };

  const handleEditEntry = (index:any) => {
    //const entryToEdit = diaries[index];
    //setNewEntry(entryToEdit);
    setEditingIndex(index);
    setShowPopup(true);
  };

  return (
    <div className="diary-app">
      <div className="diary-entries">
        {diaries.map((diary, index) => (
          <div key={index}>
            <DiaryEntry {...diary} />
            <button onClick={() => handleEditEntry(index)}>✏️</button>
          </div>
        ))}
      </div>
      {showPopup && (
        <div className="popup">
          <input type="date" name="date" value={newEntry.date} onChange={handleInputChange} />
          <input type="text" name="title" placeholder="Title" value={newEntry.title} onChange={handleInputChange} />
          <textarea name="content" placeholder="Content" value={newEntry.content} onChange={handleInputChange} />
          <button onClick={handleAddEntry}>{editingIndex !== null ? "Update Entry" : "Add Entry"}</button>
        </div>
      )}
      <button onClick={() => setShowPopup(true)}>+</button>
    </div>
  );
}

export default DiaryApp;