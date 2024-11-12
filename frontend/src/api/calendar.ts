import { Schedule } from "../types/types";
import { scheduleFromFetchedData } from "../utils/calendar";

const API_ENDPOINT: string = import.meta.env.VITE_API_ENDPOINT;

export async function fetchScheduleByEventId(
   eventId: string,
 ): Promise<Schedule | null> {
   const response = await fetch(`${API_ENDPOINT}/api/calendar/event/${eventId}`);
   const data = await response.json();
   const schedule = scheduleFromFetchedData(data)
   return schedule;
 }
 