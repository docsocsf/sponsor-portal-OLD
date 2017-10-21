import React from 'react';
import Dropzone from 'react-dropzone';
import mime from 'mime-db';
import Progress from 'Components/Progress';

const initialState = {
  loading: false,
  percent: 0,
  files: [],
  error: undefined,
};

export default class FileUploadDialog extends React.Component {
  constructor(props) {
    super(props)

    const files = props.files;
    this.state = files ? {...initialState, files} : initialState;
    this.getState = this.getState.bind(this);
    this.fileUpload = this.fileUpload.bind(this);
    this.updateFile = this.updateFile.bind(this);
  }

  async fileUpload(accepted, rejected) {
    if (accepted.length >= 1) {
      try {
        this.setState({loading: true})
        let data = await this.props.onUpload(accepted, percent =>
          this.setState({percent}));
        this.updateFile(accepted);
      } catch (e) {
        console.log(e)
        let err = <p className="error-message">{e.message}</p>
        this.setState({...initialState, error: err})
      }
    } else if (rejected.length >= 0) {
      let err = typeError(this.props.accept)
      this.setState({...initialState, error: err})
    }
  }

  typeError(valid) {
    return (
      <p className="error-message">File must be of type: {
        mime[valid].extensions.map((ext, i) =>
          <span key={i} className="extension">{ i > 0 && ", "}{ext}</span>
        )
      }</p>
    )
  }

  getState() {
    let {files, error, loading} = this.state;
    switch (true) {
      case loading: return "loading";
      case files.length > 0: return "accept";
      case !!error: return "reject";
      default: return "empty";
    }
  }

  updateFile(accepted) {
      this.setState({...initialState, files: accepted})
  }

  render() {
    let { accept, className, multiple } = this.props;
    let { files, error, loading, percent } = this.state;
    let state = this.getState();

    return (
          <Dropzone
            accept={accept}
            multiple={multiple}
            onDrop={this.fileUpload}
            className={`${state} ${className}`}
            ref={(node) => { this.dropzoneRef = node }}
          >
            <SuccessState files={files} loading={loading}/>
            <ErrorState error={error} mime={accept}/>
            <LoadingState loading={loading} percent={percent} />
            <EmptyState files={files} loading={loading} open={e => this.dropzoneRef.open()} />
          </Dropzone>
    );
  }
}

function SuccessState(props) {
  if (!props.files.length > 0 || props.loading) {
    return null;
  }

  return (
    <p>
      {props.files.map((file, i) =>
        <span key={i}>{ i > 0 && (<br/>)}{file.name}</span>
      )}
    </p>
  );
}

function ErrorState(props) {
  if (!props.error) {
    return null;
  }

  return props.error;
}

function LoadingState(props) {
  if (!props.loading) {
    return null;
  }

  return (
    <div>
      <p>
        Loading...
      </p>
      <Progress percent={props.percent} />
    </div>
  );
}

function EmptyState(props) {
  if (props.files.length > 0 || props.loading) {
    return null;
  }

  return (
    <p>
      Drag a file here or <a href="#" onClick={e => e.preventDefault()}>browse</a>.
    </p>
  );
}
