import {Question} from './questions';

export class Mucho {
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Question: Question;
  public questionid: number;
  public text: string;
  public nextQuestion: number;
  public secondToNext: string;
  public typeOfNextQuestion: string;
  public bracket: string;
}

export class Answers {
  public choices: Array<Mucho>;
}

export class AnswerResponse {
  public choice: Mucho;
}

export class MultipleChoiceAnswer {
  public question: Question;
  public text: string;
  public email: string;
}
