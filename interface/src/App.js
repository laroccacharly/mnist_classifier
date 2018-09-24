import React, { Component } from 'react';
import 'semantic-ui-css/semantic.min.css';
import {Button, Grid, Input, Table, Image, Header, Card} from "semantic-ui-react"
import { connect } from "react-redux"


let LoadSamplesButton = ({doAction}) => {
    const handleClick = () => fetch("/samples").then(data => data.json()).then(doAction).catch(console.log)
    return (
        <Button children="load samples" onClick={handleClick}/>
    )
}
LoadSamplesButton = connect(null, {doAction: (s) =>({type: "GET_SAMPLES", samples: s})})(LoadSamplesButton)

let TrainButton = () => {
    return (
        <Button children="Train"/>
    )
}

let ClassifyButton = () => {
    return (
        <Button children="Classify !"/>
    )
}

let ShowStateButton = ({state}) => {
    return <Button children="Show state" onClick={() => console.log(state)}/>
}
ShowStateButton = connect(state => ({state: state}))(ShowStateButton)

let TrainingResults = () => {
    return (
        <Card>
            <Card.Content>
                <Card.Header>Training Results</Card.Header>
                <Card.Description>Total time = 2 seconds</Card.Description>
                <Card.Description>Accuracy = 95 %</Card.Description>
            </Card.Content>

        </Card>
    )
}


let NumberImagesInput = () => {
    return (
        <Input placeholder="Number of images to train on"/>
    )
}


let SampleTableRowElement = ({sample}) => {
    return (
        <Table.Row>
            <Table.Cell>
                <Image src={sample.encodedImage}/>
            </Table.Cell>
            <Table.Cell>
                {sample.label}
            </Table.Cell>
        </Table.Row>
    )
}


let SamplesTable = ({samples}) => {

    const sampleElements = samples.map(sample => <SampleTableRowElement key={sample.id} sample={sample}/>)

    return (
        <Table basic='very' celled collapsing>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell>Images</Table.HeaderCell>
                    <Table.HeaderCell>Labels</Table.HeaderCell>
                </Table.Row>
            </Table.Header>

            <Table.Body>
                {sampleElements}
            </Table.Body>
        </Table>
    )
}
SamplesTable = connect(state => ({samples: state.samples}))(SamplesTable)


let SelectedImage = ({sample}) => {

    return (
        <div>
            <Header children="Selected Image"/>
            <Image src={sample.encodedImage}/>
        </div>
    )
}
SelectedImage = connect(state => ({sample: state.samples[0]}))(SelectedImage)



let Layout = () => {
    return (
        <Grid container >
            <Grid.Row columns={3} centered>
                <Grid.Column>
                    <div>
                        <LoadSamplesButton/>
                        <ShowStateButton/>
                    </div>
                </Grid.Column>
                <Grid.Column>
                     <div>
                         <TrainButton/>
                         <NumberImagesInput/>
                     </div>
                </Grid.Column>
                <Grid.Column>
                    <TrainingResults/>
                </Grid.Column>
            </Grid.Row>

            <Grid.Row columns={3} >
                <Grid.Column>
                    <SamplesTable/>
                </Grid.Column>
                <Grid.Column>
                    <SelectedImage/>
                </Grid.Column>
                <Grid.Column>
                    <ClassifyButton/>
                </Grid.Column>
            </Grid.Row>
        </Grid>
    )
}



class App extends Component {
  render() {
    return (
        <Layout/>
    );
  }
}

export default App;
