const actions = ['preview', 'raw']

const actionLabels = new Map(Object.entries({
    preview: 'Preview',
    raw: 'Raw',
}))

document.addEventListener('click', (e) => {
    if (!e.target.matches('button[data-action]')) {
        return;
    }
    handleAction(e);
})

function handleAction(e) {
    const actionBtn = e.target;
    const { action } = actionBtn.dataset;
    const nextAction = getNextItem(action, actions);

    // Synchronize all action buttons
    document.querySelectorAll('button[data-action]').forEach(btn => {
        btn.textContent = actionLabels.get(nextAction);
        btn.dataset.action = nextAction;
    });

    // Render the markdown
    const content = document.querySelector('textarea[name="body"]').value.trim();
    const container = document.querySelector('.rendered-markdown');
    container.innerHTML = marked.parse(content);

    // Set the next view
    document.querySelectorAll('.page-editor-view').forEach(view => {
        if (view.dataset.view === action) {
            view.dataset.active = 'true';
        } else {
            delete view.dataset.active;
        }
    });
}

function getNextItem(item, options) {
    const currentIndex = options.indexOf(item);
    const nextIndex = (currentIndex + 1) % options.length;
    return options[nextIndex];
}
