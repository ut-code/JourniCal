import { useCallback, useEffect, useState } from "react";

const API_ENDPOINT = "http://localhost:3000";

interface Entry {
  date: string;
  title: string;
  content: string;
}

export default function useJournal() {
  const [journals, setJournals] = useState<Entry[] | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const fetchEntries = useCallback(async () => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_ENDPOINT}/api/journals/`);
      if (!response.ok) {
        throw new Error("Failed to fetch journal entries");
      }
      const data = await response.json();
      setJournals(data);
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
        setJournals(null);
      }
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchEntries();
  }, [fetchEntries]);
  return { journals, isLoading, error, fetchEntries };
}
