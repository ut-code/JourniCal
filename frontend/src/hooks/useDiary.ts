import { useCallback, useEffect, useState } from "react";

const API_ENDPOINT = "http://localhost:3000";

interface Entry {
  date: string;
  title: string;
  content: string;
}

export default function useDiary() {
  const [diaries, setDiaries] = useState<Entry[] | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const fetchEntries = useCallback(async () => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_ENDPOINT}/api/diaries/`);
      if (!response.ok) {
        throw new Error("Failed to fetch diary entries");
      }
      const data = await response.json();
      setDiaries(data);
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
        setDiaries(null);
      }
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchEntries();
  }, [fetchEntries]);
  return { diaries, isLoading, error, fetchEntries };
}
