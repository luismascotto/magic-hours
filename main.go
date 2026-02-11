package main

import (
	"fmt"
	"math"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Sincronicidade struct {
	Horario  string
	Mensagem string
}

var Sincronicidades = []Sincronicidade{
	{Horario: "00:00", Mensagem: "CONEXÃO ESPIRITUAL, OBSERVE O CHAMADO DIVINO, POTENCIAL ENERGÉTICO, POSSIBILIDADES DE AUTOCONHECIMENTO"},
	{Horario: "01:01", Mensagem: "INÍCIO, NOVOS PROJETOS, CORAGEM PARA O NOVO, BUSQUE AUTOCONFIANÇA"},
	{Horario: "02:02", Mensagem: "BUSQUE SOCIALIZAR, APOSTE EM NOVAS RELAÇÕES (AMIZADES E AMOROSAS), EXPRESSE SUAS EMOÇÕES"},
	{Horario: "03:03", Mensagem: "EQUILIBRE SUAS ENERGIAS, BUSQUE ALINHAMENTO MENTAL, BUSQUE HARMONIA E PAZ, COMUNIQUE-SE"},
	{Horario: "04:04", Mensagem: "EVITE EXCESSO DE PREOCUPAÇÕES, MENOS IMPORTÂNCIA MATERIAL, ORGANIZE SUAS TAREFAS, CUIDE DO CORPO E DA MENTE"},
	{Horario: "05:05", Mensagem: "SAIA DA ZONA DE CONFORTO E DA TIMIDEZ, EXPRESSE-SE, ESCOLHA SUAS PRIORIDADES"},
	{Horario: "06:06", Mensagem: "DÊ LIMITES AOS FAMILIARES E RESGUARDE E RESPEITE ALGUMAS INTIMIDADES, ACOLHA SUA CARÊNCIA EMOCIONAL"},
	{Horario: "07:07", Mensagem: "RECONHEÇA O PONTO DA SUA VIDA QUE PRECISA DE ATENÇÃO (FINANCEIRO, ESPIRITUAL, EMOCIONAL), BUSQUE CONHECIMENTO E EVOLUÇÃO"},
	{Horario: "08:08", Mensagem: "ORGANIZE SUAS FINANÇAS, TRABALHE A PROSPERIDADE EM SUA VIDA, BUSQUE ORGANIZAR SEUS GASTOS"},
	{Horario: "09:09", Mensagem: "COLOQUE EM PRÁTICA SEUS PROJETOS INACABADOS, EVITE PENDÊNCIAS. DESAPEGUE-SE DE ROUPAS E OBJETOS QUE NÃO UTILIZA"},
	{Horario: "10:10", Mensagem: "ABANDONE SENTIMENTOS DO PASSADO, ABRA MÃO DO QUE TE PRENDE, EVOLUA COM OS ENSINOS, VIVA O PRESENTE"},
	{Horario: "11:11", Mensagem: "CONVITE PARA CONEXÃO ESPIRITUAL E AUTOCONHECIMENTO, BUSQUE FORMAS DE EVOLUÇÃO E EXPANDIR A CONSCIÊNCIA"},
	{Horario: "12:12", Mensagem: "BUSQUE EQUILIBRAR O CORPO FÍSICO, MENTAL E ESPIRITUAL. BUSQUE CONEXÃO COM A NATUREZA, COM OS DEUSES E COM AS FORÇAS QUE TE NUTREM. RELAXAMENTO E MEDITAÇÃO"},
	{Horario: "13:13", Mensagem: "RENOVE-SE, BUSQUE NOVOS HOBBIES E NOVAS METAS PARA SAIR DO COMODISMO"},
	{Horario: "14:14", Mensagem: "ALERTA DO UNIVERSO PARA SAIR DO CASULO E IR VIVER. RESPIRAR AO AR LIVRE, SORRIR E AGRADECER, CONHECER PESSOAS"},
	{Horario: "15:15", Mensagem: "NÃO DÊ IMPORTÂNCIA PARA O QUE AS PESSOAS PENSAM SOBRE VOCÊ, CONFIE NA SUA INTUIÇÃO E LIBERTE-SE, NÃO PRECISA AGRADAR A TODOS"},
	{Horario: "16:16", Mensagem: "SILENCIE, MEDITE. OUÇA, ESTUDE. SEJA RESILIENTE, MEDITE E ENCONTRE EVOLUÇÃO PESSOAL"},
	{Horario: "17:17", Mensagem: "VALORIZE O QUE REALMENTE POSSUI IMPORTÂNCIA, DÊ ATENÇÃO PARA A SUA SAÚDE, SUAS RELAÇÕES E SUA FELICIDADE"},
	{Horario: "18:18", Mensagem: "DESAPEGUE-SE DO QUE NÃO TE FAZ BEM. ABRA MÃO DE RELAÇÕES TÓXICAS, LIMPE À SUA VOLTA, LIVRE-SE DO QUE NÃO UTILIZA"},
	{Horario: "19:19", Mensagem: "REFLITA SOBRE SUA MISSÃO DE VIDA, PENSE QUAL PAPEL VOCÊ QUER DESEMPENHAR NA TERRA E O QUE VOCÊ ESTÁ FAZENDO PARA SER ALGUÉM MELHOR"},
	{Horario: "20:20", Mensagem: "HORA DE AGIR, NADA VAI CAMINHAR SEM ESFORÇO. O QUE FALTA PARA ALCANÇAR OS SEUS OBJETIVOS? O QUE VOCÊ ESTÁ FAZENDO PARA DAR CERTO?"},
	{Horario: "21:21", Mensagem: "ENSINE E AJUDE AO PRÓXIMO, SEJA A ENCONTRAR O CAMINHO DE LUZ E EVOLUÇÃO, AÇÕES SOCIAIS E AUXÍLIO À CARIDADE."},
	{Horario: "22:22", Mensagem: "DÊ ATENÇÃO PARA A SAÚDE, OBSERVE OS SINAIS QUE O SEU CORPO ESTÁ DANDO. RESGUARDE-SE, ALIMENTE-SE BEM, EXERCITE-SE E LIBERTE-SE DE VÍCIOS EMOCIONAIS"},
	{Horario: "23:23", Mensagem: "ACREDITE NAS SUAS POSSIBILIDADES, CONFIE NO SEU POTENCIAL, VALORIZE SEUS ESFORÇOS E RESPEITE O CAMINHO QUE VOCÊ VEM TRILHANDO, CONQUISTE O MUNDO."},
}

// --- Bubble Tea TUI ---
type tickMsg struct{}

type model struct {
	currentTime time.Time
	nearest     *Sincronicidade
	diffMinutes int
	shouldQuit  bool
	lastHorario string
	comboIndex  int
	termWidth   int
}

func initialModel() model {
	mrand.Seed(time.Now().UnixNano())
	now := time.Now()
	nearest, diff := findNearestSincronicidade(now)
	return model{
		currentTime: now,
		nearest:     nearest,
		diffMinutes: diff,
		lastHorario: func() string {
			if nearest != nil {
				return nearest.Horario
			}
			return ""
		}(),
		comboIndex: mrand.Intn(len(colorCombos)),
	}
}

func findNearestSincronicidade(now time.Time) (*Sincronicidade, int) {
	minutesDiff := 1 << 30
	var closest *Sincronicidade
	for i := range Sincronicidades {
		hour, _ := strconv.Atoi(Sincronicidades[i].Horario[0:2])
		minute, _ := strconv.Atoi(Sincronicidades[i].Horario[3:5])
		target := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		diff := int(math.Abs(now.Sub(target).Minutes()))
		if diff < minutesDiff {
			minutesDiff = diff
			closest = &Sincronicidades[i]
		}
	}
	return closest, minutesDiff
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg { return tickMsg{} })
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			m.comboIndex = (m.comboIndex + 1) % len(colorCombos)
			return m, nil
		case "q", "esc", "ctrl+c":
			m.shouldQuit = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
	case tickMsg:
		m.currentTime = time.Now()
		nearest, diff := findNearestSincronicidade(m.currentTime)
		m.nearest = nearest
		m.diffMinutes = diff
		if m.nearest != nil && m.nearest.Horario != m.lastHorario {
			m.lastHorario = m.nearest.Horario
			m.comboIndex = mrand.Intn(len(colorCombos))
		}
		return m, tea.Tick(time.Second, func(time.Time) tea.Msg { return tickMsg{} })
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(headerStyle.Render(" MagicHours "))
	b.WriteString("\n")
	b.WriteString(dividerStyle.Render("────────────"))
	b.WriteString("\n\n")

	combo := colorCombos[m.comboIndex]
	panel := lipgloss.NewStyle().Foreground(lipgloss.Color(combo.fg)).Background(lipgloss.Color(combo.bg)).Padding(1, 2).Bold(true)

	wrapWidth := max(int(float64(m.termWidth)*0.5), 20)
	// if m.nearest != nil && m.diffMinutes <= 5 {
	// 	wrapped := messageStyle.Width(wrapWidth).Render(m.nearest.Mensagem)
	// 	content := timeStyle.Render(m.nearest.Horario) + "\n" + wrapped
	// 	b.WriteString(panel.Render(content) + "\n\n")
	// } else if m.nearest != nil {
	if m.nearest != nil {
		if m.diffMinutes > 5 {
			b.WriteString(secondaryStyle.Render(fmt.Sprintf("Próxima: %s (em %d min)", m.nearest.Horario, m.diffMinutes)))
			b.WriteString("\n\n")
		}
		// wrapped := messageStyle.Width(wrapWidth).Render(m.nearest.Mensagem)
		// content := timeStyle.Render(m.nearest.Horario) + "\n" + wrapped
		// b.WriteString(panel.Render(content))
		b.WriteString(panel.Render(fmt.Sprintf("%s\n%s",	timeStyle.Render(m.nearest.Horario), messageStyle.Width(wrapWidth).Render(m.nearest.Mensagem))))
	} else {
		b.WriteString(secondaryStyle.Render("Nenhuma sincronicidade próxima."))
	}
	b.WriteString("\n\n")

	b.WriteString(hintStyle.Render(" q/esc para sair • c para trocar cores "))
	b.WriteString("\n")
	return b.String()
}

// --- Styling ---
type colorCombo struct {
	fg string
	bg string
}

var colorCombos = []colorCombo{
	{fg: "#FFD166", bg: "#073B4C"}, // golden on deep teal
	{fg: "#06D6A0", bg: "#1B1F3B"}, // mint on midnight
	{fg: "#EF476F", bg: "#2F2E41"}, // pink on ink
	{fg: "#A78BFA", bg: "#111827"}, // violet on near-black
	{fg: "#F59E0B", bg: "#0F172A"}, // amber on slate
}

var (
	headerStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#93C5FD")).Background(lipgloss.Color("#1F2937")).Bold(true).Padding(0, 1)
	dividerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#374151"))
	timeStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
	messageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	secondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF"))
	hintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Italic(true)
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("erro:", err)
	}
}
