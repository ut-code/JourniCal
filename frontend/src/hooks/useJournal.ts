import { useCallback, useEffect, useState } from "react";
import { Journal } from "../types/types";
import { add } from "date-fns";

const API_ENDPOINT = "http://localhost:3000";

export default function useJournal() {
  const [journals, setJournals] = useState<Journal[] | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const fetchEntries = useCallback(async () => {
    const today = new Date(new Date("2024-06-12").toDateString());
    setIsLoading(true);
    setError(null);
    const startUnixTime = Math.floor(add(today, { days: -4 }).getTime() / 1000);
    const endUnixTime = Math.floor(add(today, { days: 4 }).getTime() / 1000);

    try {
      const response = await fetch(
        `${API_ENDPOINT}/api/journals/in-range/${startUnixTime}/${endUnixTime}`,
      );
      if (!response.ok) {
        throw new Error("Failed to fetch journal entries");
      }
      const data: Journal[] = await response.json();
      const sortedData = data.sort(
        (data1, data2) =>
          new Date(data1.date).getTime() - new Date(data2.date).getTime(),
      );
      setJournals(sortedData);
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

  const fetchMoreEntriesAfter = useCallback(async (bottomDate: Date) => {
    setError(null);
    const startUnixTime = Math.floor(
      add(bottomDate, { days: 1 }).getTime() / 1000,
    );
    const endUnixTime = Math.floor(
      add(bottomDate, { days: 4 }).getTime() / 1000,
    );

    try {
      const response = await fetch(
        `${API_ENDPOINT}/api/journals/in-range/${startUnixTime}/${endUnixTime}`,
      );
      if (!response.ok) {
        throw new Error("Failed to fetch journal entries");
      }
      const data = await response.json();
      setJournals((prev) =>
        prev === null ? data : data === null ? prev : [...prev, ...data],
      );
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
        setJournals(null);
      }
    }
  }, []);

  const fetchMoreEntriesBefore = useCallback(async (topDate: Date) => {
    setError(null);
    const startUnixTime = Math.floor(
      add(topDate, { days: -3 }).getTime() / 1000,
    );
    const endUnixTime = Math.floor(topDate.getTime() / 1000);

    try {
      const response = await fetch(
        `${API_ENDPOINT}/api/journals/in-range/${startUnixTime}/${endUnixTime}`,
      );
      if (!response.ok) {
        throw new Error("Failed to fetch journal entries");
      }
      const data = await response.json();
      setJournals((prev) =>
        prev === null ? data : data === null ? prev : [...data, ...prev],
      );
    } catch (error) {
      if (error instanceof Error) {
        setError(error);
        setJournals(null);
      }
    }
  }, []);
  return {
    journals,
    isLoading,
    error,
    fetchEntries,
    fetchMoreEntriesAfter,
    fetchMoreEntriesBefore,
  };
}
