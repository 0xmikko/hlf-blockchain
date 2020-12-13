import {RouteComponentProps} from "react-router";

export interface MatchParams {
    id: string;
    tab?: string;
}

export interface DetailsScreenProps extends RouteComponentProps<MatchParams> {}
