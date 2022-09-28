export interface Player {
  id: string;
  name: string;
}

export interface DialogMessage {
  time: string;
  from: string;
  content: string;
}

export interface Room {
  id: string;
  dialogMessages: DialogMessage[];
  players: Player[];
  host: Player;
}
