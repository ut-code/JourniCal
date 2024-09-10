import { useCallback, useEffect, useState } from "react";
import { Journal } from "../types/types";

const API_ENDPOINT = "http://localhost:3000";

export default function useJournal() {
  const [journals, setJournals] = useState<Journal[] | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const fetchJournals = useCallback(async () => {
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

  const createJournal = useCallback(async (journal: Omit<Journal, "id">) => {
    try {
      const response = await fetch(`${API_ENDPOINT}/api/journals/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(journal),
      });
      if (!response.ok) {
        throw new Error("Failed to create journal");
      }
      const data = await response.json();
      setJournals((prevJournals) =>
        prevJournals ? [...prevJournals, data] : [data],
      );
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
      }
    }
  }, []);

  const updateJournal = useCallback(async (journal: Journal) => {
    try {
      const response = await fetch(`${API_ENDPOINT}/api/journals/${journal.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(journal),
      });
      if (!response.ok) {
        throw new Error("Failed to update journal");
      }
      const data = await response.json();
      setJournals((prevJournals) =>
        prevJournals
          ? prevJournals.map((j) => (j.id === data.id ? data : j))
          : [data],
      );
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
      }
    }
  }, []);

  useEffect(() => {
    fetchJournals();
  }, [fetchJournals]);
  return { journals, isLoading, error, fetchJournals, createJournal, updateJournal };
}
