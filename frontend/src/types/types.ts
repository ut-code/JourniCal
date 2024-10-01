export type Schedule = {
  id: string;
  isAllDay: boolean;
  start: Date;
  end: Date;
  title: string;
  color: string;
};

export type Journal = {
  id: number;
  date: string;
  title: string;
  content: string;
  eventId: string;
};
