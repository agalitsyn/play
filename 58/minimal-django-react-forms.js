class MyForm extends React.Component {
    constructor() {
        super();
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit(event) {
        event.preventDefault();
        const data = new FormData(event.target);

        fetch("/api/form-submit-url", {
            method: "POST",
            body: data,
        });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <CSRFToken />
                <label htmlFor="username">Enter username</label>
                <input id="username" name="username" type="text" />

                <label htmlFor="email">Enter your email</label>
                <input id="email" name="email" type="email" />

                <button>Send data!</button>
            </form>
        );
    }
}

var csrftoken = getCookie("csrftoken");

const CSRFToken = () => {
    return <input type="hidden" name="csrfmiddlewaretoken" value={csrftoken} />;
};

// https://docs.djangoproject.com/en/2.2/ref/csrf/#ajax
function getCookie(name) {
    var cookieValue = null;
    if (document.cookie && document.cookie !== "") {
        var cookies = document.cookie.split(";");
        for (var i = 0; i < cookies.length; i++) {
            var cookie = cookies[i].trim();
            // Does this cookie string begin with the name we want?
            if (cookie.substring(0, name.length + 1) === name + "=") {
                cookieValue = decodeURIComponent(
                    cookie.substring(name.length + 1)
                );
                break;
            }
        }
    }
    return cookieValue;
}
