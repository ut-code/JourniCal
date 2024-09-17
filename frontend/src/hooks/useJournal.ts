import { useCallback, useEffect, useState } from "react";
import { Journal } from "../types/types";

const API_ENDPOINT = "http://localhost:3000";

export default function useJournal() {
  const [journals, setJournals] = useState<Journal[] | null>(null);
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
