import Quill from "quill";
import Toolbar from "quill/modules/toolbar";

export class CssEditorModule {
    private toolbar: Toolbar;

    constructor(quill: Quill, options: any) {
        this.toolbar = quill.getModule('toolbar') as Toolbar;
        const toolbarContainer = this.toolbar.container;

        const buttonContainer = this.createElem("span");
        buttonContainer.setAttribute("class", "ql-formats");

        const button = this.createElem("button") as HTMLButtonElement;
        button.innerHTML = "CSS";
        button.title = "Edit page's CSS";
        button.type = "button";

        button.addEventListener("click", (event) => {
            event.preventDefault();
            this.openCssEditor();
        });

        buttonContainer.append(button);
        toolbarContainer?.append(buttonContainer);
    }

    createElem(elem: string): HTMLElement {
        return document.createElement(elem);
    }

    openCssEditor() {
        const popupContainer = this.createElem("div");
        const overlayContainer = this.createElem("div");

        const msg = "Edit HTML here, when you click \"OK\" the quill editor's contents will be replaced";

        overlayContainer.setAttribute("class", "ql-html-overlayContainer");
        popupContainer.setAttribute("class", "ql-html-popupContainer");

        const popupTitle = this.createElem("span");
        popupTitle.setAttribute("class", "ql-html-popupTitle");
        popupTitle.innerText = msg;

        const textContainer = this.createElem("div");
        textContainer.appendChild(popupTitle);
        textContainer.setAttribute("class", "ql-html-textContainer");

        const colorInput = document.createElement('input');
        colorInput.type = 'color';
        colorInput.addEventListener('input', (event) => {
            const color = (event.target as HTMLInputElement).value;
            this.applyBackgroundColor(color);
        });
        colorInput.click();

        const buttonCancel = this.createElem("button");
        buttonCancel.innerHTML = "Cancel";
        buttonCancel.setAttribute("class", "ql-html-buttonCancel");
        // const buttonOk = this.createElem("button");
        // buttonOk.innerHTML = "Ok";
        // buttonOk.setAttribute("class", "ql-html-buttonOk");
        const buttonGroup = this.createElem("div");
        buttonGroup.setAttribute("class", "ql-html-buttonGroup");

        buttonGroup.appendChild(buttonCancel);
        // buttonGroup.appendChild(buttonOk);
        textContainer.appendChild(colorInput);
        textContainer.appendChild(buttonGroup);
        popupContainer.appendChild(textContainer);
        overlayContainer.appendChild(popupContainer);

        buttonCancel.addEventListener("click", () => {
            document.body.removeChild(overlayContainer);
        });

        popupContainer.addEventListener("click", (event) => {
            event.preventDefault();
            event.stopPropagation();
        });
    }

    applyBackgroundColor(color: string) {
        const headerElement = document.querySelector('header');

        if (headerElement) {
            headerElement.style.backgroundColor = color;
        }
    }
}