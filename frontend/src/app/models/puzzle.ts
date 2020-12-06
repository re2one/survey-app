import {Question} from './questions';

export class Puzzlepiece {

  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Question: Question;
  public QuestionId: number;
  public Email: string;
  public Position: string;
  public Image: string;
  public Tapped: boolean;
}
