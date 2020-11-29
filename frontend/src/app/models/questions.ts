import {Survey} from './survey';

export class Question{
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Survey: Survey;
  public surveyid: number;
  public title: string;
  public text: string;
  public type: string;
  public first: string;
  public bracket: string;
  public next: string;
  public secondToNext: string;
  public typeOfNextQuestion: string;
}

export class Questions {
  public questions: Array<Question>;
}

export class QuestionsResponse {
  public question: Question;
}

export class FullQuestion {
  public questionId: number;
  public title: string;
  public type: string;
  public answered: boolean;
}

export class FullQuestions {
  public questions: Array<FullQuestion>;
  public finished: boolean;
}
