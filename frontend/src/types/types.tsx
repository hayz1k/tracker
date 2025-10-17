export type Track = {
  order: string;
  track: string;
  receiver: string;
  domain: string;
  adress?: string;
  note?: string;
};


export type ApiT = {
  domain: string;
  key: string;
  value: string;
  note?: string;
}