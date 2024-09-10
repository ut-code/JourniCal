import { Journal } from "../types/types";

const API_ENDPOINT: string = import.meta.env.VITE_API_ENDPOINT;

export async function fetchJournalByEventId(eventId: Journal["eventId"]): Promise<Journal | null> {
  const response = await fetch(`${API_ENDPOINT}/api/journals/event/${eventId}`);
  const data = await response.json();
  return data;
}
